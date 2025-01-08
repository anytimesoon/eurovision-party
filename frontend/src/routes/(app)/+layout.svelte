<script lang="ts">
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import MenuButton from "$lib/components/buttons/MenuButton.svelte";
  import {onMount} from "svelte";
  import {userStore} from "$lib/stores/user.store";
  import {errorStore} from "$lib/stores/error.store";
  import Toaster from "$lib/components/Toaster.svelte";
  import {CountryModel} from "$lib/models/classes/country.model";
  import {countryEP, userEP} from "$lib/models/enums/endpoints.enum";
  import {UserModel} from "$lib/models/classes/user.model";
  import {get} from "$lib/utils/genericFetch";

  let menu:HTMLElement

    onMount(async () => {
        menu = document.getElementById("menu")
        const countries = await get(countryEP.ALL) as Array<CountryModel>
        $countryStore = countries.map((country):CountryModel => {
            return CountryModel.deserialize(country)
        })

        const users = await get(userEP.ALL)
        for (const [key, value] of Object.entries(users)) {
            const user = UserModel.deserialize(value)
            $userStore[key] = user
        }
    })

  const closeMenu = () => {
    const menu = document.getElementById("menu")
    menu.classList.remove('right-0')
    menu.classList.add('-right-[75%]')
  }

  const handleWindowClick = (e:Event) => {
    const target = e.target as HTMLElement
    if(menu.classList.contains("right-0") && !target.classList.contains("voteNav") && target !== menu){
        closeMenu()
    }
  };

</script>
<svelte:window on:click={handleWindowClick} />

{#if $errorStore !== ""}
    <Toaster />
{/if}

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
          <li class="py-2">
              <a href="/country/{country.slug}" on:click={closeMenu} class="text-[1.25rem] block">
                  {country.flag} {country.name}
              </a>
          </li>
        {/each}

        <!--
            Some user reports of the countries on the bottom not being visible in chrome on iphone.
            Added some blank space to "solve", because I don't have an iphone to replicate the issue.
        -->
        <li class="my-96"></li>
      </ul>
    </div>
  </aside>

</main>











