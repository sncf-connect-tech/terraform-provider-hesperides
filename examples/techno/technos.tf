resource "hesperides_techno" "techno1" {
  name = "techno1"
  version = "1.0.0"
  working_copy = false
}

resource "hesperides_techno" "techno2" {
  name = "techno2"
  version = "2.0.0"
  working_copy = true
}
