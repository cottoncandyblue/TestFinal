package entity

import(
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"uniqueIndex" valid:"url"`
}

func TestVideo(t *testing.T){

	g := NewGomegaWithT(t)

	t.Run("check data is valid", func(t *testing.T) {
		u := Video{
			Name: "KKKK" ,
			Url: "https://www.youtube.com/",
		}
		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestName(t *testing.T){

	g := NewGomegaWithT(t)

	t.Run("check Name be blank ", func(t *testing.T) {
		u := Video{
			Name: "",
			Url: "https://www.youtube.com/",
		}
		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))	
	})
}
//forjgo

func TestUrL(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("check Url is valid",func(t *testing.T) {

		u := Video{
			Name: "gggg",
			Url: "://www.youtubess.com/",
		}

		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Url: ://www.youtubess.com/ does not validate as url"))
	})
}