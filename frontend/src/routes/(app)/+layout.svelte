<script lang="ts">
  import {currentUser} from "$lib/stores/user.store";
  import {authLvl} from "$lib/models/enums/authLvl.enum";
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import type {LayoutData} from "./$types";

  export let data:LayoutData
  $countryStore = data.countries
</script>
  
<main>
  <nav>
    <a href="/">Chat</a>

    {#each $participatingCountryStore as country }
      <span><a href={"/country/" + country.slug}>{country.name}</a>{" "}</span>
    {/each}

    <a href="/results">Results</a>

    {#if $currentUser.authLvl === authLvl.ADMIN }
      <a href="/admin/countries">Countries</a>
      <a href="/admin/friends">Friends</a>
    {/if}
  </nav>
  
  <slot />
</main>