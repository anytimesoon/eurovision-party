<script lang="ts">
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import MenuButton from "$lib/components/buttons/MenuButton.svelte";
  import {onMount} from "svelte";
  import "./../../../node_modules/@fortawesome/fontawesome-free/css/all.css"
  import {userStore} from "$lib/stores/user.store";
  import {countrySvelteEP, staticEP, userSvelteEP} from "$lib/models/enums/endpoints.enum";
  import type {ResponseModel} from "$lib/models/classes/response.model";
  import {CountryModel} from "$lib/models/classes/country.model";
  import {redirect} from "@sveltejs/kit";
  import type {UserModel} from "$lib/models/classes/user.model";

  let menu:HTMLElement

  onMount(async () => {
      menu = document.getElementById("menu")

      if($countryStore.length === 0) {
          const countryRes = await fetch(countrySvelteEP.ALL)
          const countries: ResponseModel<CountryModel[]> = await countryRes.json()

          $countryStore = countries.body.map((country):CountryModel => {
              return new CountryModel().deserialize(country)
          })
      }

      const usersRes = await fetch(userSvelteEP.ALL)
      const users: ResponseModel<Map<string, UserModel>> = await usersRes.json()

      $userStore = users.body
  })

  const closeMenu = () => {
    const menu = document.getElementById("menu")
    menu.classList.remove('right-0')
    menu.classList.add('-right-[75%]')
  }

  const handleWindowClick = (e:Event) => {
    let el = e.target as HTMLElement
    if(menu.classList.contains("right-0") && !el.classList.contains("voteNav") && e.target !== menu){
      closeMenu()
    }
  };

</script>
<svelte:head>
        <link rel="preload" as="image" href={staticEP.IMG + "/content/static/img/newuser.png"} />
</svelte:head>
<svelte:window on:click={handleWindowClick} />

<main class="h-screen max-w-screen-sm mx-auto px-3 relative flex flex-col-reverse">

        <nav class="flex w-full items-center justify-between pb-1">

            <MenuButton navName="Chat" dest="/" menu={menu}/>
            <MenuButton navName="Vote" dest="" menu={menu}/>
            <MenuButton navName="Results" dest="/results"  menu={menu}/>
            <MenuButton navName="Settings" dest="/settings"  menu={menu}/>

        </nav>

        <div id="content" class="flex-1 pb-4">
            <slot />
        </div>



  <aside id="menu" class="fixed top-0 -right-[75%] w-[75%] bg-canvas-secondary z-1 flex duration-500 h-screen overflow-auto">
    <div class="w-full flex flex-col p-5">
      <ul class="list-none">
        {#each $participatingCountryStore as country}
          <li class="py-2"><a href="/country/{country.slug}" on:click={closeMenu} class="text-[1.25rem]">{country.flag} {country.name}</a></li>
        {/each}
      </ul>
    </div>
  </aside>

</main>











