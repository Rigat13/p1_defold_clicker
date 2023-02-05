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
    value: "3.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "autoclicker_sprite"
  type: "sprite"
  data: "tile_set: \"/main/assets/Sprites/sprite_animation_Anass.atlas\"\n"
  "default_animation: \"Anass_idle\"\n"
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
    x: 0.166
    y: 0.166
    z: 0.166
  }
}
