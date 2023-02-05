components {
  id: "autoclicker"
  component: "/main/assets/scripts/autoclicker/autoclicker.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "click_time"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "autoclicker_sprite"
  type: "sprite"
  data: "tile_set: \"/main/assets/textures/autoclicker_texture.atlas\"\n"
  "default_animation: \"autoclicker_texture\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.2
    y: 0.2
    z: 0.2
  }
}
