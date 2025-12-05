package main

import (
	"encoding/json"
	"path/filepath"
	"fmt"
	"bytes"
	"strings"
	"os"
	"os/exec"
	"crypto/rand"
	"encoding/base64"
	"log"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func getAsset32randomByteString() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func getAssetPath(videoID, mediaType string) string {
	ext := mediaTypeToExt(mediaType)
	return fmt.Sprintf("%s%s", videoID, ext)
}

func (cfg apiConfig) getAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg apiConfig) getAssetURL(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

func mediaTypeToExt(mediaType string) string {
	parts := strings.Split(mediaType, "/")
	if len(parts) != 2 {
		return ".bin"
	}
	return "." + parts[1]
}

func getVideoAspectRatio(filepath string) (string, error) {
    type parameters struct {
        Streams []struct {
            Ratio string `json:"display_aspect_ratio"`
        } `json:"streams"`
    }

    cmd := exec.Command("ffprobe",
        "-v", "error",
        "-print_format", "json",
        "-show_streams",
        filepath,
    )

    var out bytes.Buffer
    cmd.Stdout = &out

    if err := cmd.Run(); err != nil {
        log.Print("Couldn't run cmd: ", err)
        return "", err
    }

    var params parameters
    if err := json.Unmarshal(out.Bytes(), &params); err != nil {
        log.Print("Couldn't decode params: ", err)
        log.Print("Raw ffprobe output: ", out.String())
        return "", err
    }

    if len(params.Streams) == 0 {
        return "", fmt.Errorf("no streams found")
    }

    ratio := params.Streams[0].Ratio
    if ratio != "9:16" && ratio != "16:9" {
        return "other", nil
    }

    return ratio, nil
}
