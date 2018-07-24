resource "hesperides_techno" "techno1" {
  name = "techno1"
  version = "1.0.0"
  working_copy = false

  templates = [
    {
      name = "temp1"
      namespace = "temp1_namespace1"
      filename = "temp1_filename1"
      location = "temp1_location1"
      content = "temp1_content1"
      rights = {
        user = {
          read = false
          write = false
          execute = false
        }
        group = {
          read = false
          write = false
          execute = false
        }
        other = {
          read = false
          write = false
          execute = false
        }
      }
      version_id = 1
    },
  ]
}

resource "hesperides_techno" "techno2" {
  name = "techno2"
  version = "2.0.0"
  working_copy = true

  templates = [
    {
      name = "temp2"
      namespace = "temp2_namespace2"
      filename = "temp2_filename2"
      location = "temp2_location2"
      content = "temp2_content2"
      rights = {
        user = {
          read = true
          write = true
          execute = true
        }
        group = {
          read = true
          write = true
          execute = true
        }
        other = {
          read = true
          write = true
          execute = true
        }
      }
      version_id = 1
    },
  ]
}
