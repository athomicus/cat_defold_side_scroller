components {
  id: "trees_bushes"
  component: "/input/trees_bushes.tilemap"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"/input/trees_bushes.tilemap\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"trees\"\n"
  "mask: \"player\"\n"
  ""
}
