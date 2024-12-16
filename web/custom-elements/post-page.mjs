function PostPage({ html, state }) {
  const { store } = state;
  const { post } = store;
  const { frontmatter } = post;
  const {
    description = "",
    published = "",
    title = "",
    components = [],
  } = frontmatter;

  // Function to render components based on frontmatter keys
  const renderComponents = () => {
    return components
      .map((component) => {
        switch (component) {
          case "counter":
            return "<custom-counter></custom-counter>";
          case "h-card":
            return '<my-h-card class="hidden"></my-h-card>';
          // Add more component cases as needed
          default:
            return "";
        }
      })
      .join("\n");
  };

  return html`
    <style>
      h1,
      .date {
        text-align: var(--align-heading);
      }

      .post-layout {
        display: grid;
        grid-template-columns: 1fr 300px;
        gap: 2rem;
      }
    </style>
    <site-layout>
      <div class="post-layout">
        <article class="h-entry font-body leading4 mi-auto pb0 pb4-lg">
          <a href="/" class="text-link">← Back</a>

          <h1
            class="p-name font-heading font-bold mbe0 text4 tracking-1 leading1"
          >
            ${title}
          </h1>
          <p class="date dt-published mbe4">${published}</p>
          <aside class="post-components">${renderComponents()}</aside>

          <section class="prose" slot="e-content doc">${post.html}</section>
          <section class="p-summary hidden">${description}</section>
        </article>
      </div>
    </site-layout>
  `;
}

//   ${mentions?.length ? "<webmentions-list></webmentions-list>" : ""}
// <my-h-card class="hidden"></my-h-card>;
