<script lang="ts">
  import {currentUser} from "$lib/stores/user.store";
  import {authLvl} from "$lib/models/enums/authLvl.enum";
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import {CountryModel} from "$lib/models/classes/country.model";
  import {UserModel} from "$lib/models/classes/user.model";
  import type {LayoutData} from "../../../.svelte-kit/types/src/routes/$types";

  export let data:LayoutData
  $countryStore = data.countries

</script>
  
<main>
  <nav>
    <a href="/">Home</a>

    {#each $participatingCountryStore as country }
      <span><a href={"/country/" + country.slug}>{country.name}</a>{" "}</span>
    {/each}

    {#if $currentUser.authLvl === authLvl.ADMIN }
      <a href="/admin/countries">Countries</a>
      <a href="/admin/friends">Friends</a>
    {/if}
  </nav>
  
  <slot />
</main>