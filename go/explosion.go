components {
  id: "timer_message"
  component: "/scripts/timer_message.script"
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
    id: "delay"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "path"
    value: "/game_manager"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "component"
    value: "delete_sender"
    type: PROPERTY_TYPE_HASH
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/assets/sprites.atlas\"\n"
  "default_animation: \"explosion\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.5
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
