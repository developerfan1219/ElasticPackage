<template>
  <v-responsive
      class="mx-auto bg-white"            
  >
  <v-container
    class="spacing-playground py-16 bg-white service-section"
    fluid
    max-width="1400"
  >
    <v-row>
      <v-col cols="12" md="4">
        <v-card class="bg-white" variant="text">
          <v-card-item>
            <v-card-title class="service-title">For Customers</v-card-title>
            
          </v-card-item>

          <v-card-text class="service-desc">
            Helping packaging clients access industry best prices and solutions, with real time visibility over their orders.
          </v-card-text>
        </v-card>        
      </v-col>

      <v-col cols="12" md="4">
        <v-card class="bg-white" variant="text">
          <v-card-item>
            <v-card-title class="service-title">For Dealerships</v-card-title>
            
          </v-card-item>

          <v-card-text class="service-desc">
            Helping sales professionals build their own packaging distribution businesses, focused on customers, not admin.
          </v-card-text>
        </v-card>        
      </v-col>

      <v-col cols="12" md="4">
        <v-card class="bg-white" variant="text">
          <v-card-item>
            <v-card-title class="service-title">For Suppliers</v-card-title>

          </v-card-item>

          <v-card-text class="service-desc">
            Helping manufacturers and distributors win more business, by providing feedback on every quote.
          </v-card-text>
        </v-card>        
      </v-col>
    </v-row>
  </v-container>
  </v-responsive>
</template>

<style>
.service-title {
  color: #333;
  font-size: 32px;
  font-weight: 500;
}
.service-desc {
  color: #333;
  font-size: 18px;
  font-weight: 400;
}
</style>