USE
[BankAccount]
GO
/****** Object:  Table [dbo].[Accounts]    Script Date: 13/04/2022 16:24:13 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Accounts]
(
    [
    Id] [
    bigint]
    NOT
    NULL, [
    FirstName] [
    nvarchar]
(
    50
) NOT NULL,
    [LastName] [nvarchar]
(
    50
) NOT NULL,
    [IBAN] [nvarchar]
(
    20
) NOT NULL,
    [Balance] [bigint] NOT NULL,
    [CreationDate] [datetime2]
(
    7
) NOT NULL,
    CONSTRAINT [PK_Accounts_1] PRIMARY KEY CLUSTERED
(
[
    Id] ASC
) WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON)
  ON [PRIMARY]
    )
  ON [PRIMARY]
    GO
/****** Object:  Table [dbo].[TransactionHistories]    Script Date: 13/04/2022 16:24:13 ******/
    SET ANSI_NULLS
  ON
    GO
    SET QUOTED_IDENTIFIER
  ON
    GO
CREATE TABLE [dbo].[TransactionHistories]
(
    [
    Id] [
    bigint]
    NOT
    NULL, [
    SourceAccountId] [
    bigint]
    NOT
    NULL, [
    DestinationAccountId] [
    bigint]
    NOT
    NULL, [
    TransactionType] [
    int]
    NOT
    NULL, [
    Amount] [
    bigint]
    NOT
    NULL, [
    CurrentBalance] [
    bigint]
    NOT
    NULL, [
    TransactionDateTime] [
    datetime2]
(
    7
) NOT NULL,
    CONSTRAINT [PK_TransactionHistories_1] PRIMARY KEY CLUSTERED
(
[
    Id] ASC
) WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON)
  ON [PRIMARY]
    )
  ON [PRIMARY]
    GO
