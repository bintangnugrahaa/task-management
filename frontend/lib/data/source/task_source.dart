import 'dart:convert';

import 'package:frontend/common/urls.dart';
import 'package:frontend/data/models/task.dart';
import 'package:d_method/d_method.dart';
import 'package:http/http.dart' as http;
import 'package:image_picker/image_picker.dart';

class TaskSource {
  /// `'${URLs.host}/tasks'`
  static const _baseURL = '${URLs.host}/tasks';

  static Future<bool> add(
    String title,
    String description,
    String dueDate,
    int userId,
  ) async {
    try {
      final response = await http.post(
        Uri.parse(_baseURL),
        body: jsonEncode({
          "title": title,
          "description": description,
          "status": "Queue",
          "dueDate": dueDate,
          "userId": userId
        }),
      );
      DMethod.logResponse(response);

      return response.statusCode == 201;
    } catch (e) {
      DMethod.log(e.toString(), colorCode: 1);
      return false;
    }
  }
      int userId, String status) async {
    try {
      final response = await http.get(
        Uri.parse('$_baseURL/user/$userId/$status'),
      );
      DMethod.logResponse(response);

      if (response.statusCode == 200) {
        List resBody = jsonDecode(response.body);
        return resBody.map((e) => Task.fromJson(Map.from(e))).toList();
      }

      return null;
    } catch (e) {
      DMethod.log(e.toString(), colorCode: 1);
      return null;
    }
  }
}
