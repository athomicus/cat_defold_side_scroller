components {
  id: "cavern"
  component: "/input/cavern.tilemap"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"/input/cavern.tilemap\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"terrain,obstacles\"\n"
  "mask: \"player\"\n"
  ""
}
