package sub

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/hu-jinwen/Typora-img-upload/cmd"
	"github.com/spf13/cobra"
	"time"
)

var (
	akID       string
	akSecret   string
	endpoint   string
	bucketName string
	path       string
)

func init() {
	ossCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "oss endpoint（非空）")
	ossCmd.Flags().StringVarP(&bucketName, "bucket", "b", "", "oss bucketName（非空）")
	ossCmd.Flags().StringVarP(&akID, "access_id", "a", "", "oss AccessKey ID（非空）")
	ossCmd.Flags().StringVarP(&akSecret, "access_secret", "s", "", "oss AccessKey Secret（非空）")
	ossCmd.Flags().StringVarP(&path, "path", "p", "pic-bed", "oss 路径前缀，可以为空")

	cmd.RootCmd.AddCommand(ossCmd)
}

var ossCmd = &cobra.Command{
	Use:   "oss",
	Short: "Typora 图片上传 oss 插件",
	Long:  "使用oss当图床的 Typora 图片上传插件",
	RunE: func(cmd *cobra.Command, args []string) error {
		if akID == "" || akSecret == "" || endpoint == "" || bucketName == "" {
			return errors.New("akID or akSecret or endpoint or bucket must not be empty")
		}
		client, err := oss.New(endpoint, akID, akSecret)
		if err != nil {
			return err
		}
		bucket, err := client.Bucket(bucketName)
		if err != nil {
			return err
		}
		return uploadToOss(bucket, args)
	},
}

func nowDateStr() (dateStr string) {
	return time.Now().Format("2006-01-02")
}

// md5Str 字符串做md5
func md5Str(str string) (md5Str string) {
	bytes := []byte(str)
	sum := md5.Sum(bytes)
	return fmt.Sprintf("%x", sum)
}

func uploadToOss(bucket *oss.Bucket, args []string) (err error) {
	// /Users/hujinwen/Downloads/sss.png
	for _, filePath := range args {
		savePath := fmt.Sprintf("%s/%s/%s.png", path, nowDateStr(), md5Str(filePath))
		err = bucket.PutObjectFromFile(savePath, filePath)
		if err != nil {
			return
		}
		imgUrl := fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, savePath)
		fmt.Println(imgUrl)
	}
	return
}
