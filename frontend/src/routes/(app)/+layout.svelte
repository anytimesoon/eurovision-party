<script lang="ts">
  // import "../../../frontend/src/app.css";
  import {authLvl} from "$lib/models/enums/authLvl.enum";
  import {partCountryStore} from "$lib/stores/partCountry.store";
  import {currentUserStore} from "$lib/stores/user.store";
</script>

<svelte:head>
  <meta name="title" content="Eurovision party" />
  <meta name="description" content="Chat and vote in a private remote party" />
</svelte:head>
  
<main>
  <nav>
    <a href="/">Home</a>

    {#each $partCountryStore as country }
        <a href={"/country/" + country.slug}>{country.name}</a>
    {/each}

    {#if $currentUserStore != null && $currentUserStore.authLvl === authLvl.ADMIN}
      <a href="/admin/countries">Countries</a>
      <a href="/admin/friends">Friends</a>
    {/if}
  </nav>
  
  <slot />
</main>