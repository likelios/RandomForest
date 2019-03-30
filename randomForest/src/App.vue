<template>
  <v-app>
    <v-navigation-drawer
      app
      temporary
      v-model="drawer"
    >
      <v-list>

        <v-list-group
          prepend-icon="filter"
          value="true"
        >
          <v-list-tile slot="activator">
            <v-list-tile-title>Фильтр</v-list-tile-title>
          </v-list-tile>

          <v-list-group
            no-action
            sub-group
            value="true"
          >
            <v-list-tile slot="activator">
              <v-list-tile-title>Admin</v-list-tile-title>
            </v-list-tile>

            <v-list-tile

              @click=""
            >
              <v-list-tile-title></v-list-tile-title>
              <v-list-tile-action>
                <v-icon></v-icon>
              </v-list-tile-action>
            </v-list-tile>
          </v-list-group>


        </v-list-group>

        <v-list-tile
          v-for="link of links"
          :key="link.title"
          :to="link.url"
        >
          <v-list-tile-action>
            <v-icon>{{link.icon}}</v-icon>
          </v-list-tile-action>

          <v-list-tile-content>
            <v-list-tile-title v-text="link.title"></v-list-tile-title>
          </v-list-tile-content>

        </v-list-tile>


        <v-list-tile slot="activator">
          <v-list-tile-title>Users</v-list-tile-title>
        </v-list-tile>

        <v-list-tile
          @click="onLogout"
          v-if="isUserLoggedIn"
        >
          <v-list-tile-action>
            <v-icon>exit_to_app</v-icon>
          </v-list-tile-action>

          <v-list-tile-content>
            <v-list-tile-title v-text="'Выйти'"></v-list-tile-title>
          </v-list-tile-content>

        </v-list-tile>

      </v-list>
    </v-navigation-drawer>

    <v-toolbar app dark absolute clipped-left color="primary">

      <v-toolbar-side-icon
        class="hidden-md-and-up"
        @click="drawer = !drawer"
      ></v-toolbar-side-icon>

      <v-toolbar-title>
        <v-layout align-center justify-center>
          <router-link to="/" tag="span" class="pointer ml-5">
            RandomForest
          </router-link>
        </v-layout>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-sm-and-down">

        <v-btn
          v-for="link in links"
          :key="link.title"
          :to="link.url"
          flat
        >
          <v-icon left>{{link.icon}}</v-icon>
          {{link.title}}
          <v-btn color="red" fab small dark v-if="link.count">
            {{link.count}}
          </v-btn>
        </v-btn>

        <v-btn
          @click="onLogout"
          flat
          v-if="isUserLoggedIn"
        >
          <v-icon left>exit_to_app</v-icon>
          Выйти
        </v-btn>
        <v-divider
          class="mx-2"
          inset
          vertical
        ></v-divider>



      </v-toolbar-items>

    </v-toolbar>

    <v-content>
      <router-view></router-view>
    </v-content>

    <template v-if="error">

      <v-snackbar
        :multi-line="true"
        :timeout="5000"
        @input="closeError"
        @value="true"
        color="error"
      >
        {{error}}
        <v-btn
          dark
          flat
          @click.native="closeError"
        >
          Close
        </v-btn>
      </v-snackbar>
    </template>


  </v-app>

</template>

<script>
  export default {
    name: 'App',
    data() {
      return {
        drawer: false,
      }
    },
    methods: {
      closeError() {
        this.$store.dispatch('clearError')
      },
      onLogout() {
        this.$store.dispatch('logoutUser');
        this.$storage.remove('test')
        console.log(this.$storage)
        this.$router.push('/')

      }
    },
    computed: {
      error() {
        return this.$store.getters.error
      },
      isUserLoggedIn() {
        return this.$store.getters.isUserLoggedIn
      },
      links() {
        if (this.isUserLoggedIn) {
          this.$storage.set('test', {key: this.$store.getters.user.id}, {ttl: 60 * 1000})
          if (this.$store.getters.getUser.hasOwnProperty('manager_id')) {
            return [
              {title: 'Панель администрирования', icon: 'perm_identity', url: '/admin', count: ''},
            ]
          } else {
            return [
              {title: 'Личный кабинет', icon: 'perm_identity', url: '/orders', count: ''},
              {title: 'Корзина', icon: 'shopping_cart', url: '/list', count: `${this.getTovars}`},
            ]
          }

        } else {
          return [
            {title: 'Логин', icon: 'lock', url: '/login', count: ''},
            {title: 'Регистрация', icon: 'face', url: '/registration', count: ''},
          ]
        }
      }
    },
    created() {
    },
    props: {
      source: String
    }
  }
</script>

<style>

  .pointer {
    cursor: pointer;
  }

  .application {
    font-family: 'Fira Sans' !important;
  }

  .navigation-drawer__border {
    display: none;
  }


</style>
