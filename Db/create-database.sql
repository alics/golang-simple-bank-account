CREATE DATABASE [BankAccount]
GO

USE [BankAccount];
GO

CREATE TABLE Accounts (
    [Id] [bigint] NOT NULL,
    [FirstName] [nvarchar](50) NOT NULL,
    [LastName] [nvarchar](50) NOT NULL,
    [IBAN] [nvarchar](20) NOT NULL,
    [Balance] [bigint] NOT NULL,
    [CreationDate] [datetime2](7) NOT NULL,
                         PRIMARY KEY (Id)
);
GO


CREATE TABLE TransactionHistories (
    [Id] [bigint] NOT NULL,
    [SourceAccountId] [bigint] NOT NULL,
    [DestinationAccountId] [bigint] NOT NULL,
    [TransactionType] [int] NOT NULL,
    [Amount] [bigint] NOT NULL,
    [CurrentBalance] [bigint] NOT NULL,
    [TransactionDateTime] [datetime2](7) NOT NULL,
    PRIMARY KEY (Id)
    );
GO