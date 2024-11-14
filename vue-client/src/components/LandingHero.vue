<template>
  <v-responsive
      class="mx-auto bg-black"
  >
  <v-container
    class="landing-hero spacing-playground py-16 "
    fluid
    max-width="1400"
  >
  <v-row>
      
    <v-col cols="5" class="px-8 d-flex flex-column ga-sm-6 justify-center">
        <h1 class="text-h3 font-weight-bold hero-title">How we built USA's largest distribution network</h1>
        <p class="hero-desc text-h5 ">Without owning a single warehouse or a truck.</p>            
    </v-col>
    <v-col cols="7">
      <v-card class="rounded-xl">
      <div id="maplayout">
        <div id="mapsidebar">
          <h2>2024</h2>
          <div>
            <v-icon
              color="orange-darken-2"
              icon="mdi-thumb-up"
            ></v-icon>
            Suppliers
          </div>
          <div>
            <v-icon
              color="blue-lighten-2"
              icon="mdi-thumb-up"
            ></v-icon>
            Dealerships
          </div>
        </div>
        <Map v-model="location" />
      </div>
      </v-card>
    </v-col>
  </v-row>
  </v-container>
</v-responsive>
</template>

<style>  

  .hero-title {
    line-height: 69px;  
  }
  .landing-hero .hero-desc {
    line-height: 40px;
    font-size: 28px !important; 
  }
  .mapboxgl-control-container {
    display: none;
  }
  #maplayout {
    flex: 1;
    display: flex;
    height: 450px;
    position: relative;
  }

  #mapsidebar {
    background-color: #000;
    color: #fff;
    padding: 6px 12px;
    
    z-index: 1;
    position: absolute;
    bottom: 0;
    right: 0;
    margin: 12px;
    border-radius: 20px;
  }
  
</style>

<script>    

  export default {
    components: {
      Map
    },

    data: () => ({
      location: {
        lng: -97.8655,
        lat: 39.8733,
        bearing: 0,
        pitch: 0,
        zoom: 4
      }
    })
  };
</script>