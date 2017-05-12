#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    void * wrapper; // ImagePublisherWrapper, because we can't use any c++ class directly in cgo.
}ImagePublisher;

ImagePublisher* NewImagePublisher(char *ip, char *topic);
void PublishImage(ImagePublisher* pub, unsigned char*data, int len);
void DeleteImagePublisher(ImagePublisher* pub);

#ifdef __cplusplus
}
#endif
