package encrypt

import "golang.org/x/crypto/bcrypt"

// CompareHashAndPassword is
func CompareHashAndPassword(hashedPassword, password *string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(*password))
	if err != nil {
		return err
	}
	return nil
}
