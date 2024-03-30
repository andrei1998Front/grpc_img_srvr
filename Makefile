test:
	mkdir internal\storage\test_imgs_dir
	go test ./internal/storage
	rmdir /s /q internal\storage\test_imgs_dir