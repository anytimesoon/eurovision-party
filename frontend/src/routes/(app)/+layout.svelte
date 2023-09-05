<script lang="ts">
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import type {LayoutData} from "./$types";
  import MenuButton from "$lib/components/buttons/MenuButton.svelte";
  import {onMount} from "svelte";
  import "./../../../node_modules/@fortawesome/fontawesome-free/css/all.css"


  export let data:LayoutData
  $countryStore = data.countries
  let menu:HTMLElement

  onMount(() => {
    menu = document.getElementById("menu")
  })

  const closeMenu = () => {
    const menu = document.getElementById("menu")
    menu.classList.remove('right-0')
    menu.classList.add('-right-56')
  }

  const handleWindowClick = (e:Event) => {
    if(menu.classList.contains("right-0") && !e.target.classList.contains("voteNav") && e.target !== menu){
      closeMenu()
    }
  };

</script>
<svelte:window on:click={handleWindowClick} />
<main class="h-screen max-w-screen-sm mx-auto p-3 relative">
  <div class="flex flex-col">
      <div id="content" class="flex h-[calc(100vh-5rem)] pb-3">
        <div class="flex-grow">
          <slot />
        </div>
      </div>
      <nav class="h-4 flex w-full flex-wrap items-center justify-between">

            <MenuButton icon="chat" menu={menu}/>
            <MenuButton icon="votes" menu={menu}/>
            <MenuButton icon="results"  menu={menu}/>
            <MenuButton icon="settings"  menu={menu}/>

      </nav>
  </div>

  <aside id="menu" class="fixed top-0 -right-56 bg-white z-1 flex duration-500 h-screen overflow-auto">
    <div class="w-full flex flex-col p-5">
      <ul class="list-none">
        {#each $participatingCountryStore as country}
          <li class="py-2"><a href="/country/{country.slug}" on:click={closeMenu}>{country.flag} {country.name}</a></li>
        {/each}
      </ul>
    </div>
  </aside>

</main>











