package main

func getLocalName(lang string, names map[string]string) string {
	// Если мапы нет в базе (город/страна не найдены), возвращаем заглушку
	if names == nil {
		return "Unknown"
	}

	// 1. Проверяем запрошенный язык (например, "ru")
	if name, exists := names[lang]; exists && name != "" {
		return name
	}

	// 2. Откатываемся на английский, если запрошенного языка не нашлось
	if name, exists := names["en"]; exists && name != "" {
		return name
	}

	// 3. Крайний случай: если нет ни "lang", ни "en", возвращаем любой имеющийся перевод
	for _, name := range names {
		if name != "" {
			return name
		}
	}

	return "Unknown"
}
