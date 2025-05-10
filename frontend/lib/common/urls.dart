class URLs {
  static const host = 'http://192.168.80.1:8080';
  static String image(String fileName) => '$host/attachments/$fileName';
}
