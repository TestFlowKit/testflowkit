export default {
  import: ['dist/index.js'],
  paths: ['features/**/*.feature'],
  format: ['progress', 'html:reports/cucumber-report.html'],
};
