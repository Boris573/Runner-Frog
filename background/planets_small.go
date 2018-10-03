components {
  id: "planets"
  component: "/background/planets.script"
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
  id: "sprite1"
  type: "sprite"
  data: "tile_set: \"/background/background.atlas\"\n"
  "default_animation: \"ice_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 57.0
    y: 103.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: -0.30007803
    w: 0.95391464
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "tile_set: \"/background/background.atlas\"\n"
  "default_animation: \"gas_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -478.0
    y: -230.0
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
  id: "sprite3"
  type: "sprite"
  data: "tile_set: \"/background/background.atlas\"\n"
  "default_animation: \"ring_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 300.0
    y: -373.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.63373095
    w: 0.77355355
  }
}
