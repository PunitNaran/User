package user

import (
	"localhost/users"

	flatbuffers "github.com/google/flatbuffers/go"
)

func MakeUser(b *flatbuffers.Builder, user *User) []byte {
	// re-use the already-allocated Builder:
	b.Reset()

	users.UserAddSexuality(b, makeUserSexuality(b, &user.Sexuality))

	users.VerificationStart(b)
	users.VerificationAddEncryptedEmail(b, b.CreateString(user.Profile.Verification.EncryptedEmail))
	users.VerificationAddEncryptedPass(b, b.CreateString(user.Profile.Verification.EncryptedPass))
	users.VerificationAddEncryptedMobileNo(b, b.CreateString(user.Profile.Verification.EncryptedMobileNo))
	users.UserAddAge(b, makeUserSexuality(b, &user.Sexuality))
	users.UserAddReligon(b, makeUserRelgion(b, &user.Religon))
	users.VerificationAddIncorrectCounts(b, user.Profile.Verification.IncorrectCounts)
}

func makeUserSexuality(b *flatbuffers.Builder, s *Sexuality) flatbuffers.UOffsetT {
	return users.CreateSexuality(b, s.Heterosexual, s.Homosexual, s.Bisexual)
}

func makeUserAge(b *flatbuffers.Builder, a *Age) flatbuffers.UOffsetT {
	return users.CreateAge(b, a.Day, a.Month, a.Year)
}

func makeUserRelgion(b *flatbuffers.Builder, r *Religon) flatbuffers.UOffsetT {
	users.BuddhismAddCommunity(b, b.CreateString(r.Buddhism.Community))

	users.ChristianAddCommunity(b, b.CreateString(r.Christian.Community))

	users.HinduismAddCommunity(b, b.CreateString(r.Hinduism.Community))

	users.IslamAddCommunity(b, b.CreateString(r.Islam.Community))

}
