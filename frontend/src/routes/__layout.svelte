<script>
  import "../app.css";
  export const trailingSlash = 'ignore';
</script>

<svelte:head>
  <meta name="title" content="Eurovision party" />
  <meta name="description" content="Chat and vote in a private remote party" />
</svelte:head>
  
<main>
  <nav>
    <a href="/">Home</a>
    <a href="/countries">Countries</a>
    <a href="/friends">Friends</a>
  </nav>
  
  <slot />
</main>