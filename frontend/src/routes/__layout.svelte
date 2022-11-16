<script>
    import "../app.css";
  </script>
  
  <svelte:head>
    <meta name="title" content="Eurovision party" />
    <meta name="description" content="Chat and vote in a private remote party" />
  </svelte:head>
  
  <nav>
    <a href=".">HOME</a>
    <a href="/about">ABOUT</a>
  </nav>
  
  <slot />
  
  <style>
    nav {
      padding: 1rem;
    }
    a {
      text-decoration: none;
      color: gray;
      margin-right: 1rem;
    }
  </style>
  