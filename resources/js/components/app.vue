<template lang="pug">
  .app.app-header-fixed(:class="folded == true ? 'app-aside-folded': ''")
    AppHeaderComponent(v-show="authenticated")
    AppAsideComponent(:folded="folded" v-on:header-folded="setFolded" v-show="authenticated")

    #content(role="main" :class="authenticated ? 'app-content': ''")
      .app-content-body.app-content-full
        .hbox.hbox-auto-xs.hbox-auto-sm.bg-light
          router-view
    AppFooterComponent(v-show="authenticated")
</template>

<script>
import { event } from '@/services/bus'
import { mapGetters } from 'vuex'
export default {
  name: 'AppComponent',
  data () {
    return {
      appReady: false,
      folded: false
    }
  },
  computed: {
    ...mapGetters({
      authenticated: 'IsUserAllowSeeNavigator'
    })
  },
  components: {
    AppHeaderComponent: () => import('./header.vue'),
    AppAsideComponent: () => import('./aside.vue'),
    AppFooterComponent: () => import('./footer.vue')
  },

  methods: {
    setFolded (payload) {
      this.folded = !this.folded
    }
  }
}
</script>

