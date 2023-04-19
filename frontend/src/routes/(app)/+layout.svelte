<script lang="ts">
  import "../../../frontend/src/app.css";
  import {currentUser} from "$lib/helpers/user.helper.ts";
  import {authLvl} from "$lib/models/enums/authLvl.enum.ts";
  import {partCountryStore} from "$lib/stores/partCountry.store.js";
  export const trailingSlash = 'ignore';
  const me = currentUser();
</script>

<svelte:head>
  <meta name="title" content="Eurovision party" />
  <meta name="description" content="Chat and vote in a private remote party" />
</svelte:head>
  
<main>
  <nav>
    <a href="/frontend/static">Home</a>

    {#each $partCountryStore as country }
        <a href={"/country/" + country.slug}>{country.name}</a>
<!--      <p>hello</p>-->
    {/each}

    {#if me != null && me.authLvl === authLvl.ADMIN}
      <a href="/admin/countries">Countries</a>
      <a href="/admin/friends">Friends</a>
    {/if}
  </nav>
  
  <slot />
</main>