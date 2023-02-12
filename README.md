Go cache service

usage example:


	cache := cache.New()

	cache.Set("userName", "username")
	username, ok := cache.Get("userName")
	fmt.Println(username, ok)

	ok = cache.Delete("userId")

	username, ok = cache.Get("userName")
	fmt.Println(username, ok)