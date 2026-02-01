@FILE_UPLOAD @FORM @FRONTEND
Feature: File Upload e2e tests

    Background:
        Given the user goes to the "file upload e2e" page

    @SINGLE_FILE_UPLOAD
    Scenario: a user can upload a single file successfully
        When the user uploads the "avatar_image" file into the "Avatar" field
        Then the "files uploaded block" should contain the text "avatar.png"

    @MULTIPLE_FILE_UPLOAD
    Scenario: a user can upload multiple files successfully
        When the user uploads the "gallery_image1, gallery_image2, gallery_image3" files into the "Gallery" field
        Then the "files uploaded block" should contain the text "image1.jpg"
        And the "files uploaded block" should contain the text "image2.jpg"
        And the "files uploaded block" should contain the text "image3.jpg"

    @DOCUMENT_UPLOAD
    Scenario: a user can upload a document successfully
        When the user uploads the "test_document" file into the "Document" field
        Then the "files uploaded block" should contain the text "test.pdf"

    @CSV_UPLOAD
    Scenario: a user can upload a CSV file successfully
        When the user uploads the "sample_csv" file into the "Data" field
        Then the "files uploaded block" should contain the text "sample.csv"


