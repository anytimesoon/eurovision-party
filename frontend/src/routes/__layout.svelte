<script>
  import "../app.css";
  import { getCountries } from "$lib/stores/country.story";
  import Nav from "$lib/components/Nav.svelte"
  export const trailingSlash = 'ignore';
  getCountries();
</script>

<svelte:head>
  <meta name="title" content="Eurovision party" />
  <meta name="description" content="Chat and vote in a private remote party" />
</svelte:head>
  
<main>
  <Nav />
  <slot />
</main>