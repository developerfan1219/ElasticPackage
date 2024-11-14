<template>
  <v-container
      class="mx-auto"
      max-width="1400"
      width="100%"
      
    >
    
    <v-app-bar  :elevation="2" class="app-header position-sticky px-3 bg-black justify-center">
        

        <v-app-bar-title class="text-h2 site-title font-weight-bold">
          <RouterLink to="/" >Packform</RouterLink>
          
        </v-app-bar-title>

        <div class="d-flex justify-content ga-sm-6">
        
        <v-list-item link to="/orders" class="text-h5 menu-title font-weight-bold" title="Orders"></v-list-item>
        <v-list-item link to="/teams" class="text-h5 menu-title font-weight-bold" title="Team"></v-list-item>
        <v-list-item link to="/media" class="text-h5 menu-title font-weight-bold" title="Media"></v-list-item>
        </div>

        <template v-slot:append>
          <div class="d-flex justify-content ga-sm-3">
            <v-btn variant="outlined" class="sign-up-btn">Sign up</v-btn>
            <v-btn variant="elevated" class="login-btn">Login</v-btn>            
          </div>
        </template>

    </v-app-bar>

  </v-container>
    
</template>

<style>
.site-title {
  line-height: initial;
  
}
.site-title a {
  text-decoration: none;
  color: #fff;
}
.app-header {
  height: 118px;

}
.app-header .v-toolbar-title {
  line-height: initial !important;
}

.app-header .menu-title .v-list-item-title {
  font-size: 24px !important; 
  font-weight: 700;
}

.app-header .sign-up-btn {
  border-color: #fe8002;
  color: #fe8002;
  font-weight: 700;
  font-size: 18px;
  text-transform: unset;
}
.app-header .sign-up-btn:hover {
  background-color: #e5eafd;
}
.app-header .login-btn {
  background-color: #fe8002;
  font-weight: 700;
}
</style>

<script setup>
  import { ref } from 'vue'

  const theme = ref('dark')

  function onClick () {
    theme.value = theme.value === 'light' ? 'dark' : 'light'
  }
</script>