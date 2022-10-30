components {
  id: "move"
  component: "/scripts/move.script"
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
    id: "speed"
    value: "1000.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "trigger_responce_delete"
  component: "/scripts/trigger_responce.script"
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
    id: "mask"
    value: "enemy"
    type: PROPERTY_TYPE_HASH
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
components {
  id: "trigger_responce_explosion"
  component: "/scripts/trigger_responce.script"
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
    id: "mask"
    value: "enemy"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "path"
    value: "/spawner"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "component"
    value: "explosion_spawner"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "message"
    value: "spawn"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "trigger_responce_explosion_small_enemy"
  component: "/scripts/trigger_responce.script"
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
    id: "mask"
    value: "small_enemy"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "path"
    value: "/spawner"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "component"
    value: "explosion_spawner"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "message"
    value: "spawn"
    type: PROPERTY_TYPE_HASH
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/assets/sprites.atlas\"\n"
  "default_animation: \"bullet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: -0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"bullet\"\n"
  "mask: \"enemy\"\n"
  "mask: \"small_enemy\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 13.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 6.281\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  "bullet: false\n"
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
}
