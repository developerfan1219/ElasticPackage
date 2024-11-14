<template>
  <v-responsive
      class="mx-auto bg-white"
      
      
  >
  <v-container
    class="spacing-playground py-16 bg-white action-section"
    fluid
    max-width="1400"
  >

    <div>
      <div class="d-flex justify-space-between ga-sm-8">
        <v-btn variant="elevated" class="action-btn btn-green">Customer</v-btn>
        <v-btn variant="elevated" class="action-btn btn-blue">Dealership</v-btn>            
        <v-btn variant="elevated" class="action-btn btn-orange">Supplier</v-btn>            
      </div>
    </div>    
  
  </v-container>
  </v-responsive>
</template>

<style>
.action-btn {
  font-size: 32px;
  padding: 12px 24px;
  min-width: 354px;
  height: 100% !important;
  text-transform: inherit;
}
.action-btn.btn-green {
  background-color: #28a745;
}
.action-btn.btn-blue {
  background-color: #007bff;
}
.action-btn.btn-orange {
  background-color: #fe8002;
}
.action-section {
  border-bottom: 1px solid #ddd;
}
</style>