<template>
    <v-responsive
      class="mx-auto bg-white"            
  >
  <v-container
    class="spacing-playground py-16 bg-white text-section"
    fluid
    max-width="1400"
  >
    <v-row>
      <v-col cols="3">
        <h1 class="text-h3 content-title">What is Packform?</h1>
      </v-col>

      <v-col cols="9" class="d-flex flex-column ga-sm-4">
        <div class="content-unit">
          Packform is the <span class="content-emphas">technology connecting suppliers, dealers and customers</span>, for USA's largest Packaging distribution network.
        </div>
        <div class="content-unit">
          The platform streamlines the packaging procurement process, enabling Packform dealerships to provide <span class="content-emphas">quick quotes through its global supplier network</span>, 
          while allowing customers to <span class="content-emphas">easily approve, manage, and track deliveries</span>.
        </div>
        <div class="content-unit">
          Powering the USA's <span class="content-emphas">largest B2B distribution network</span> (without owning a single truck or warehouse), Packform is transforming the way businesses procure goods.
        </div>
      </v-col>
    </v-row>
  </v-container>
  </v-responsive>
</template>

<style>
.text-section {
  border-bottom: 1px solid #ddd;
}
.text-section .content-title {
  font-size: 64px !important;
  font-weight: 500;
  line-height: 92px;
  color: #333;
}

.content-unit {
  font-size: 20px;
  font-weight: 400;
  color: #333;
}
.content-emphas {
  color: #fe8002;
}
</style>