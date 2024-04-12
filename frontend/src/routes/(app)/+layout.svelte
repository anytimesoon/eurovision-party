<script lang="ts">
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import type {LayoutData} from "./$types";
  import MenuButton from "$lib/components/buttons/MenuButton.svelte";
  import {onMount} from "svelte";
  import {userStore} from "$lib/stores/user.store";


  export let data:LayoutData

  let menu:HTMLElement

  onMount(() => {
    menu = document.getElementById("menu")
  })

  const closeMenu = () => {
    const menu = document.getElementById("menu")
    menu.classList.remove('right-0')
    menu.classList.add('-right-[75%]')
  }

  const handleWindowClick = (e:Event) => {
    if(menu.classList.contains("right-0") && !e.target.classList.contains("voteNav") && e.target !== menu){
      closeMenu()
    }
  };

  $: if (data.countries) {
      $countryStore = data.countries
  }

  $: if (data.users) {
      $userStore = data.users
  }
</script>
<svelte:window on:click={handleWindowClick} />

<main class="h-screen max-w-screen-sm mx-auto px-3 relative flex flex-col">
    <div id="content" class="flex-1 pb-4 overflow-hidden">
        <slot />
    </div>

    <nav class="flex w-full items-center justify-between pb-1">
        <MenuButton navName="Chat" dest="/" menu={menu}/>
        <MenuButton navName="Vote" dest="" menu={menu}/>
        <MenuButton navName="Results" dest="/results"  menu={menu}/>
        <MenuButton navName="Settings" dest="/settings"  menu={menu}/>
    </nav>




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











