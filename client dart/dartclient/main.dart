import 'dart:io';
import 'dart:convert';

void main() async {
  final certFilePath = 'cert/server-cert.pem';
  try {
    final certBytes = await File(certFilePath).readAsBytes();
    final securityContext = SecurityContext(withTrustedRoots: true);
    securityContext.setTrustedCertificatesBytes(certBytes);

    final client = HttpClient(context: securityContext);
    final request = await client.getUrl(Uri.parse('https://localhost:1150/api/hello'));

    final response = await request.close();
    final responseBody = await response.transform(utf8.decoder).join();

    print('Response: $responseBody');
  } catch (e) {
    print('Failed: $e');
  }
}
