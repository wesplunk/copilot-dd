import spock.lang.Specification

class PhotoServiceSpec extends Specification {
    def "should add a photo and return the assigned id"() {
        given:
        PhotoService photoService = new PhotoService()
        Photo photo = new Photo()

        when:
        int id = photoService.addPhoto(photo)

        then:
        id == 0
    }

    def "should get a photo by id"() {
        given:
        PhotoService photoService = new PhotoService()
        Photo photo = new Photo()
        int id = photoService.addPhoto(photo)

        when:
        Photo result = photoService.getPhoto(id)

        then:
        result == photo
    }

    def "should delete a photo by id"() {
        given:
        PhotoService photoService = new PhotoService()
        Photo photo = new Photo()
        int id = photoService.addPhoto(photo)

        when:
        photoService.deletePhoto(id)

        then:
        photoService.getPhoto(id) == null
    }

    def "should update a photo by id"() {
        given:
        PhotoService photoService = new PhotoService()
        Photo photo = new Photo()
        int id = photoService.addPhoto(photo)
        Photo updatedPhoto = new Photo()

        when:
        photoService.updatePhoto(id, updatedPhoto)

        then:
        photoService.getPhoto(id) == updatedPhoto
    }
}