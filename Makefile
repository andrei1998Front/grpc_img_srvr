test:
	mkdir internal\storage\test_imgs_dir
	go test ./internal/storage -cover
	rmdir /s /q internal\storage\test_imgs_dir