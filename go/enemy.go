components {
  id: "trigger_responce_bullet"
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
    value: "ship"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "path"
    value: "/ship"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "component"
    value: "ship_controller"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "message"
    value: "hit"
    type: PROPERTY_TYPE_HASH
  }
}
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
}
components {
  id: "auto_rotate"
  component: "/scripts/auto_rotate.script"
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
    value: "10.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "trigger_responce_gc"
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
    value: "gc"
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
  id: "enemy"
  component: "/scripts/enemy.script"
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
    id: "hp"
    value: "10.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "damage_group"
    value: "bullet"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "add_score_final_message"
  component: "/scripts/add_score_final_message.script"
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
    id: "path"
    value: "/game_manager"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "component"
    value: "game"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "score"
    value: "10.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "final_message"
  component: "/scripts/final_message.script"
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
    id: "path"
    value: "/spawner"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "component"
    value: "final_explosion_spawner"
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
  "default_animation: \"roid-large\"\n"
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
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"ship\"\n"
  "mask: \"bullet\"\n"
  "mask: \"gc\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: 3.0\n"
  "      y: -17.0\n"
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
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: -10.0\n"
  "      y: 20.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 1\n"
  "    count: 1\n"
  "  }\n"
  "  data: 63.9625\n"
  "  data: 59.15\n"
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
