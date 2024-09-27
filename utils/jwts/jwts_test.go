package jwts

import "testing"

func TestGenerateJwtToken(t *testing.T) {
	token, err := GenerateToken(JwtPayload{
		UserId:   1,
		NickName: "fsdafsda",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
