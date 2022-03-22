package replace

import (
	"log"
	"strings"
)

func Replace(original_str string) string {
	var new_string string

	re := strings.NewReplacer("Oracle", "Oracle©", "Google", "Google©", "Microsoft", "Microsoft©", "Amazon", "Amazon©", "Deloitte", "Deloitte©")
	new_string = re.Replace(original_str)
	log.Println("Replaced string from function ", new_string)
	return new_string
}
