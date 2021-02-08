<!DOCTYPE html>
<!--
Copyright 2014 Mozilla Foundation
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->
<html dir="ltr" mozdisallowselectionprint>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <meta name="google" content="notranslate">
  <title>PDF.js page viewer using built components</title>

  <style>
    body {
      background-color: #808080;
      margin: 0;
      padding: 0;
    }
  </style>

  <link rel="stylesheet" href="/static/js/pdfjs/web/pdf_viewer.css">

  <script src="/static/pdfjs/pdf.js"></script>
  <script src="/static/pdfjs-dist/web/pdf_viewer.js"></script>
</head>

<body tabindex="1">
  <div id="pageContainer" class="pdfViewer singlePageView"></div>

  <script src="/static/js/pdfjs/pageviewer.js"></script>
</body>
</html>