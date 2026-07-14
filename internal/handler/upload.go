package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func saveUploadedFile(c *fiber.Ctx, fieldName, destSubdir string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil || file == nil {
		return "", nil
	}

	safeName := sanitizeFilename(file.Filename)
	if safeName == "" {
		safeName = "upload"
	}

	ext := strings.ToLower(filepath.Ext(safeName))
	if ext == "" {
		ext = filepath.Ext(file.Filename)
	}
	if ext == "" {
		ext = ".bin"
	}

	uploadDir, err := absoluteUploadDir(destSubdir)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d-%s%s", time.Now().UnixNano(), strings.TrimSuffix(safeName, ext), ext)
	filename = strings.ReplaceAll(filename, " ", "_")
	filename = strings.Trim(filename, "-._")

	destPath := filepath.Join(uploadDir, filename)
	if err := c.SaveFile(file, destPath); err != nil {
		return "", err
	}

	return "/uploads/" + strings.Trim(destSubdir, "/") + "/" + filename, nil
}

func saveUploadedFiles(c *fiber.Ctx, fieldName, destSubdir string) ([]string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, nil
	}
	if form == nil || len(form.File[fieldName]) == 0 {
		return nil, nil
	}

	uploadDir, err := absoluteUploadDir(destSubdir)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		return nil, err
	}

	var urls []string
	for _, file := range form.File[fieldName] {
		safeName := sanitizeFilename(file.Filename)
		if safeName == "" {
			safeName = "upload"
		}

		ext := strings.ToLower(filepath.Ext(safeName))
		if ext == "" {
			ext = filepath.Ext(file.Filename)
		}
		if ext == "" {
			ext = ".bin"
		}

		filename := fmt.Sprintf("%d-%s%s", time.Now().UnixNano(), strings.TrimSuffix(safeName, ext), ext)
		filename = strings.ReplaceAll(filename, " ", "_")
		filename = strings.Trim(filename, "-._")

		destPath := filepath.Join(uploadDir, filename)
		if err := saveMultipartFile(file, destPath); err != nil {
			return nil, err
		}

		urls = append(urls, "/uploads/"+strings.Trim(destSubdir, "/")+"/"+filename)
	}

	return urls, nil
}

func saveMultipartFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func absoluteUploadDir(destSubdir string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	uploadDir := filepath.Join(cwd, "storage", "uploads", destSubdir)
	return uploadDir, nil
}

func sanitizeFilename(name string) string {
	name = filepath.Base(name)
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, "..", "")
	allowed := regexp.MustCompile(`[^a-zA-Z0-9._-]+`)
	return allowed.ReplaceAllString(name, "-")
}
