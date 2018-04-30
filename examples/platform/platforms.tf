resource "hesperides_platform" "platform_USN1" {
  application = "${hesperides_application.application_app1.name}"
  name = "USN1"
  version = "1.0.0"
  production = false
}

resource "hesperides_platform" "platform_USN2" {
  application = "${hesperides_application.application_app1.name}"
  name = "USN2"
  version = "1.0.0"
  production = false
}

resource "hesperides_platform" "platform_INT1" {
  application = "${hesperides_application.application_app1.name}"
  name = "INT1"
  version = "1.0.0"
  production = false
}

resource "hesperides_platform" "platform_INT2" {
  application = "${hesperides_application.application_app1.name}"
  name = "INT2"
  version = "1.0.0"
  production = false
}
