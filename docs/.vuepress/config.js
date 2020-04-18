module.exports = {
  base: "/gokit/",
  title: "gokit",
  themeConfig: {
    repo: "bongnv/gokit",
    editLinks: true,
    sidebar: [{
      title: "Guide",
      collapsable: false,
      sidebarDepth: 2,
      children: [
        '',
        'guide',
        'scaffolding-projects',
        'scaffolding-crud-endpoints',
        'generating-services',
        'generating-dao'
      ]
    }]
  }
}