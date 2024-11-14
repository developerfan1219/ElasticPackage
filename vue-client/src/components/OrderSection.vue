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
      
    <v-col cols="12" class="px-8 d-flex flex-column ga-sm-6 justify-center">
      <v-card
        title="Orders"
        flat
      >
        <template v-slot:text>
          <v-text-field
            v-model="search"
            label="Search"
            prepend-inner-icon="mdi-magnify"
            variant="outlined"
            hide-details
            single-line
          ></v-text-field>
        </template>

        <h4 class="ml-4">Created Date</h4>
        <div class="d-flex ml-4">
          
          <v-date-input
            v-model="selectedDates"
            label="Select range"
            max-width="368"
            multiple="range"
          ></v-date-input>
        </div>

        <div v-if="serverItems.length > 0" class="d-flex ml-4">
          <h4>Total Amount: ${{ totalPrice }}</h4>
        </div>

        <v-data-table-server
          v-model:items-per-page="itemsPerPage"
          :headers="headers"          
          :items="serverItems"
          :items-length="totalItems"
          :loading="loading"
          :search="search"          
          item-value="name"
          @update:options="loadItems"
        ></v-data-table-server>
      </v-card>        
    </v-col>
  </v-row>
  </v-container>
</v-responsive>
</template>

<script>
  import axios from 'axios'

  import { en } from 'vuetify/locale';
  const api_url = import.meta.env.VITE_APP_API_BASE_URL;
  
  export default {
    data: () => ({
      totalPrice: 0,
      pageNum: 1,
      itemsPerPage: 5,
      selectedDates: null,
      start: null,
      end: null,      
      headers: [
        {
          title: 'Order name',
          align: 'start',
          sortable: false,
          key: 'order_name',
        },
        { title: 'Customer Company', key: 'company_name', align: 'end' },
        { title: 'Customer Name', key: 'user_id', align: 'end' },
        { title: 'Order Date', key: 'created_at', align: 'end' },
        { title: 'Delivered Amount', key: 'delivered_amount', align: 'end' },
        { title: 'Total Amount', key: 'total_quantity', align: 'end' },
      ],
      search: '',
      serverItems: [],
      loading: true,
      totalItems: 0,
    }),
    watch: {
      async selectedDates(newValue, oldValue) {

        if (newValue && newValue.length > 1) {
          let start = this.formatDate(newValue[0]),
              end   = this.formatDate(newValue[newValue.length - 1], false);          
          this.start = start;
          this.end = end;
        }
      },
      start: 'fetchData',
      end: 'fetchData',
      serverItems: 'calculateTotalAmount'
    },
    methods: {
      calculateTotalAmount() {
        let total = 0;
        this.serverItems.forEach(item => {
          total += parseInt(item.delivered_amount || 0);
        });
        this.totalPrice = total;
      },
      formatDate(date, start=true) {
          let d = new Date(date),
              month = '' + (d.getMonth() + 1),
              day = '' + d.getDate(),
              year = d.getFullYear();              

          if (month.length < 2) 
              month = '0' + month;
          if (day.length < 2) 
              day = '0' + day;
          if (start)
            return [year, month, day].join('-')+'T'+'00:00:00Z';
          else
            return [year, month, day].join('-')+'T'+'23:59:59Z';
      },
      async fetchData() {
        this.loading = true
        try {
          console.log(this.pageNum, this.itemsPerPage, this.start, this.end);
          const response = await axios.post(`${api_url}/order_details`, {     
            search: this.search,
            start: this.start || '',       
            end: this.end || '',
            page: this.pageNum,
            count: this.itemsPerPage            
          });
          
          if (response.data.items != null) {
            this.serverItems = response.data.items;
          } else {
            this.serverItems = [];
          }
          

          this.totalItems = response.data.total;
          this.loading = false
        } catch (error) {
          console.error("Error fetching orders:", error);
          this.loading = false
        }
      },
      async loadItems ({ page, itemsPerPage}) {
        this.pageNum = page
        this.itemsPerPage = itemsPerPage
        this.fetchData()
      },
    },
  }
</script>