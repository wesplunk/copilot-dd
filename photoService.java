# write photo service to handle crud on photo ids
class PhotoService {
    private Map<Integer, Photo> photoMap = new HashMap<>();
    private int id = 0;

    public int addPhoto(Photo photo) {
        photoMap.put(id, photo);
        return id++;
    }

    public Photo getPhoto(int id) {
        return photoMap.get(id);
    }

    public void deletePhoto(int id) {
        photoMap.remove(id);
    }

    public void updatePhoto(int id, Photo photo) {
        photoMap.put(id, photo);
    }
}
