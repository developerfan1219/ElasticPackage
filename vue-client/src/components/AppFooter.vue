<template>  
  <v-container 
  class="spacing-playground service-section"
  fluid
  max-width="1400"
  width="100%"
  >
  <div class="d-flex flex-column align-start ga-sm-3">
    <div class="d-flex footer-main-section">
      <div class="footer-logo">
        Packform
      </div>
      <div class="d-flex footer-link-section">
        <div class="link-item">
          <div>Packform</div>
          <a>About Us</a>
          <a>Team</a>
          <a>Media</a>
        </div>
        <div class="link-item">
          <div>Spend Less</div>
          <a>Customers</a>
          <a>Find Dealership</a>
        </div>
        <div class="link-item">
          <div>Earn More</div>
          <a>Suppliers</a>
          <a>Dealers</a>
        </div>
        <div class="link-item">
          <div>Follow Us</div>
          <div>
            <a
              v-for="item in items"
              :key="item.title"
              :href="item.href"
              :title="item.title"
              class="d-inline-block mx-2 social-link"
              rel="noopener noreferrer"
              target="_blank"
            >
              <v-icon
                :icon="item.icon"
                :size="32"
              />
            </a>
          </div>
        </div>
      </div>
    </div>
    
    <div class="copyrigt-section">
      &copy; 2024-{{ (new Date()).getFullYear() }} <span class="d-none d-sm-inline-block">Packform Pty Ltd</span>
      —
      <a
        class="on-surface"
        href="https://vuetifyjs.com/about/licensing/"
        rel="noopener noreferrer"
        target="_blank"
      >
      Privacy Policy
      </a>
      <a
        class="on-surface"
        href="https://vuetifyjs.com/about/licensing/"
        rel="noopener noreferrer"
        target="_blank"
      >
      Terms of Use
      </a>
      
    </div>
      
  </div>  
  </v-container>

</template>

<style>
.copyrigt-section a {
  text-decoration: underline;
  margin-left: 8px;
}
.copyrigt-section a::before {
  content: " ";
    width: 6px;
    height: 6px;
    background: #fff;
    display: inline-block;
    border-radius: 3px;
    margin-bottom: 4px;
}
.footer-main-section {
  width: 100%;
}
.link-item {
  flex: 1 0 0;
  display: flex;
  flex-direction: column;
  row-gap: 12px;
}
.link-item a {
  text-decoration: underline;
}
.footer-logo {
  flex: 0 0 30%;
  font-size: 48px;
  font-weight: 700;
}
.footer-link-section {
  flex: 0 0 70%;
}
</style>

<script setup lang="ts">
  const items = [
    {
      title: 'Vuetify Documentation',
      icon: `$vuetify`,
      href: 'https://vuetifyjs.com/',
    },
    {
      title: 'Vuetify Support',
      icon: 'mdi-shield-star-outline',
      href: 'https://support.vuetifyjs.com/',
    },
    {
      title: 'Vuetify X',
      icon: ['M2.04875 3.00002L9.77052 13.3248L1.99998 21.7192H3.74882L10.5519 14.3697L16.0486 21.7192H22L13.8437 10.8137L21.0765 3.00002H19.3277L13.0624 9.76874L8.0001 3.00002H2.04875ZM4.62054 4.28821H7.35461L19.4278 20.4308H16.6937L4.62054 4.28821Z'],
      href: 'https://x.com/vuetifyjs',
    },
    {
      title: 'Vuetify GitHub',
      icon: `mdi-github`,
      href: 'https://github.com/vuetifyjs/vuetify',
    },
    
  ]
</script>

<style scoped lang="sass">
  .social-link :deep(.v-icon)
    color: #FFF
    text-decoration: none
    transition: .2s ease-in-out    

    &:hover
      color: rgba(25, 118, 210, 1)
</style>
