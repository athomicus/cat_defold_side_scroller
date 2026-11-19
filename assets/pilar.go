embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"pilar_dol\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/input/cavern.tilesource\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"pilar_middle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 32.0\n"
  "  y: 32.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/input/cavern.tilesource\"\n"
  "}\n"
  ""
  position {
    y: 32.0
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "default_animation: \"pilar_up\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 32.0\n"
  "  y: 32.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/input/cavern.tilesource\"\n"
  "}\n"
  ""
  position {
    y: 64.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"obstacles\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 20.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 11.457284\n"
  "  data: 30.458809\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
