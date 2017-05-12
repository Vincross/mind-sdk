#include <rosskill.h>
#include <ros.h>
#include <embedded_linux_hardware.h>
#include <sensor_msgs/CompressedImage.h>

class ImagePublisherWrapper {
    private:
        // Default OUTPUT_SIZE of NodeHandle is too small to send images, we have to create it manually.
        ros::NodeHandle_<EmbeddedLinuxHardware, 25, 25, 512, 32767> nh;
        sensor_msgs::CompressedImage image;
        ros::Publisher publisher;
    public:
        ImagePublisherWrapper(char *ip, char *topic): publisher(topic, &image) {
            nh.initNode(ip);
            nh.advertise(publisher);
            image.format = "jpeg";
        }
        void publishImage(unsigned char* data, int len) {
            image.data_length = len;
            image.data = (unsigned char*)data;
            publisher.publish(&image);
            nh.spinOnce();
        }
};

/* Wrapper functions */

extern "C" ImagePublisher* NewImagePublisher(char *ip, char *topic) {
    ImagePublisher* img_pub = new ImagePublisher();
    img_pub->wrapper = new ImagePublisherWrapper(ip, topic);
    return img_pub;
}

extern "C" void PublishImage(ImagePublisher* img_pub, unsigned char*data, int len) {
    ((ImagePublisherWrapper *)img_pub->wrapper)->publishImage(data, len);
}

extern "C" void DeleteImagePublisher(ImagePublisher* img_pub) {
    ImagePublisherWrapper* wrapper = (ImagePublisherWrapper *)(img_pub->wrapper);
    delete wrapper;
    delete img_pub;
}
