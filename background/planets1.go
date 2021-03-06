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
    x: -138.0
    y: -265.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.011050336
    w: 0.99993896
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
    x: -449.0
    y: 97.0
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
    x: 87.0
    y: 121.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: -0.17540981
    w: 0.9844955
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/background/background.atlas\"\n"
  "default_animation: \"earthlike_planet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 707.0
    y: -40.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
