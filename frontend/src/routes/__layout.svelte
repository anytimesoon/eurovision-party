<script>
  import "../app.css";
  import {currentUser} from "$lib/helpers/user.helper.ts";
  import {authLvl} from "$lib/models/enums/authLvl.enum.ts";
  export const trailingSlash = 'ignore';
  const me = currentUser();
</script>

<svelte:head>
  <meta name="title" content="Eurovision party" />
  <meta name="description" content="Chat and vote in a private remote party" />
</svelte:head>
  
<main>
  <nav>
    <a href="/">Home</a>
    {#if me != null && me.authLvl === authLvl.ADMIN}
      <a href="/countries">Countries</a>
      <a href="/friends">Friends</a>
    {/if}
  </nav>
  
  <slot />
</main>